const core =  require('@actions/core')
// const github = require('@actions/github') // provide access to github rest api
const exec = require('@actions/exec')

function run(){
    const bucket = core.getInput('buckets',{ required: true});
    const bucket_region = core.getInput('buckets-region',{required: true});
    const deploy = core.getInput('dist-folder',{required:true});


    core.notice('Hello from deploy to AWS S3');
    const s3Uri = `s3://${bucket}`
    exec.exec(`aws s3 sync ${deploy} ${s3Uri} --region ${bucket_region} `)
    
    const websiteUrl= `http://${bucket}.s3-website.${bucket_region}.amazonaws.com`
    core.setOutput('website-url',websiteUrl)

}

run();