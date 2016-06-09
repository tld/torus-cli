'use strict';

var childProcess = require('child_process');

var _ = require('lodash');
var Promise = require('es6-promise').Promise;
var regulator = require('event-regulator');

var validate = require('../validate');
var cValue = require('./value');
var credentials = require('./credentials');

var run = exports;

var validator = validate.build({
  environment: validate.slug,
  service: validate.slug,
  instance: validate.slug
});

run.execute = function (ctx) {
  return new Promise(function (resolve, reject) {
    var params = ctx.params;
    var serviceName = ctx.option('service').value;
    var envName = ctx.option('environment').value;
    var instance = ctx.option('instance').value;

    if (params.length < 1) {
      return reject(new Error('You must provide commands to run'));
    }

    if (!serviceName) {
      return reject(new Error('You must provide a --service flag'));
    }

    if (!envName) {
      return reject(new Error('You must provide a --environment flag'));
    }

    var errors = validator({
      service: serviceName,
      environment: envName,
      instance: instance
    });
    if (errors.length > 0) {
      return reject(errors[0]);
    }

    return credentials.get(ctx.session, {
      service: serviceName,
      environment: envName,
      instance: instance
    }).then(function (creds) {
      return run.spawn(ctx.daemon, params, creds);
    }).catch(reject);
  });
};

run.spawn = function (daemon, params, creds) {
  return new Promise(function (resolve, reject) {
    var env = _.clone(process.env);
    creds.forEach(function (cred) {
      var value = cValue.parse(cred.body.value);
      if (value.body.type === 'undefined') {
        return;
      }

      env[cred.body.name.toUpperCase()] = value.body.value;
    });

    var proc = childProcess.spawn(params[0], params.slice(1), {
      cwd: process.cwd(),
      detached: false,
      env: env
    });

    proc.stdout.pipe(process.stdout);
    proc.stderr.pipe(process.stderr);

    function onClose(exitCode) {
      // If the exitCode is 0 then propagate up a successful run to
      // `bin/arigatoo` which will exit with a non-zero exitCode.
      //
      // TODO: Don't swallow the exitCode and propagate it all the way up.
      daemon.disconnect().then(function () {
        resolve(exitCode === 0);
      }).catch(function (err) {
        console.error('Could not disconnect from daemon');
        reject(err);
      });
    }

    function handleSignal(signal) {
      console.log('sending signal', signal);
      proc.kill(signal);
    }

    regulator.subscribe([
      [process, 'SIGHUP', handleSignal.bind(this, 'SIGHUP')],
      [process, 'SIGINT', handleSignal.bind(this, 'SIGINT')],
      [process, 'SIGQUIT', handleSignal.bind(this, 'SIGQUIT')],
      [process, 'SIGTERM', handleSignal.bind(this, 'SIGTERM')],

      [proc, 'error', reject, true],
      [proc, 'close', onClose, true]
    ]);
  });
};