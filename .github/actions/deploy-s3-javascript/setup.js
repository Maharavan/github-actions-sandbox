const core =  require('@actions/core')
const github = require('@actions/exec')
const exec = require('@actions/exec')

async function setup(){
    core.notice('Upload file begins');
}

setup();