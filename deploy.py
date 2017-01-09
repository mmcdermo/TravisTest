import red_leader
import argparse

def deploy_prod():
    rl = red_leader.RedLeader(
        cluster_prefix="signal-api",
        image_name="289307572179.dkr.ecr.us-west-1.amazonaws.com/signal_api_server",
        aws_profile="signal",
        aws_region="us-west-1"
    )
    tag = rl._ecr_push_build("signal_api_server", ".")
    rl.update_default_service_image_tag(tag)

def deploy_dev():
    rl = red_leader.RedLeader(
        cluster_prefix="signal-api-dev",
        image_name="289307572179.dkr.ecr.us-west-1.amazonaws.com/signal_api_dev_server",
        aws_profile="signal",
        aws_region="us-west-1"
    )
    tag = rl._ecr_push_build("signal_api_dev_server", ".")
    rl.update_default_service_image_tag(tag)

def main():
    parser = argparse.ArgumentParser()
    group = parser.add_mutually_exclusive_group(required=True)
    group.add_argument("--production", help="Build and deploy to production", action="store_true")
    group.add_argument("--dev", help="Build and deploy to dev servers", action="store_true")
    args = parser.parse_args()

    if args.production is True:
        print("Deploying for production...")
        deploy_prod()
    elif args.dev is True:
        print("Deploying for development...")
        deploy_dev()
    else:
        raise Exception("Neither production nor dev flags set")
    
if __name__ == '__main__':
    main()
