## Start unbuntu18 compile docker container
`docker run -it --network host --rm --cap-add=SYS_PTRACE --name ub18 -v /Users/zzy/Documents/agora/media_build:/home/jenkins/media_build -v /Users/zzy/Documents/agora/sync:/home/jenkins/sync hub.agoralab.co/jenkins/ubuntu/18_04/m64:compile-ubuntu-mix-streaming`