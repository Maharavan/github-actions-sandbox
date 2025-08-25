const core = require("@actions/core")
const github = require("@actions/github")
const exec = require("@actions/exec")

async function post(){
    core.notice('Upload file begins');
}

post();