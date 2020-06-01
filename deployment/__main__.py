import pulumi
import pulumi_gcp as gcp

def CloudRun( img="gcr.io/leadmrktr/pasteit",location="asia-northeast1",project="leadmrktr"):
    cr = gcp.cloudrun.Service("pasteit",
                location=location,
                project=project,
                template={
                    "spec": {
                        "containers": [{
                            "image": img,
                        }],
                    },
                },
                traffics=[{
                    "latestRevision": True,
                    "percent": 100,
                }],
        )
    gcp.cloudrun.IamMember("enable_for_all",service=cr.name,project= project,location=location,  role= "roles/run.invoker",member="allUsers")
    return cr


cloud_run = CloudRun()
pulumi.export("url",cloud_run.status["url"])
pulumi.export("cloud run",cloud_run)