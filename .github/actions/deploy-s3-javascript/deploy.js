const core =  require('@actions/core')
const github = require('@actions/exec')
const exec = require('@actions/exec')

function run(){
    core.notice('Hello from deploy to AWS S3');
}

run();