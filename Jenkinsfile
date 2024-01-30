pipeline {
    agent any

    parameters {
        choice(choices: ["dev", "sit", "prod"], description: "Which environment to deploy?", name: "deployEnvironment")
    }

    environment {
        GOLANG_IMAGE_NAME = "vocaverse-golang"
        CONTAINER_NAME = "vocaverse-golang"
    }

    stages {
        stage ('Remove container'){
                    steps {
                    script {
                            def exitCode = sh(script: "docker rm -f ${CONTAINER_NAME}-${params.deployEnvironment}", returnStatus: true)

                            if (exitCode == 0) {
                                echo "Container removal was successful"
                            } else {
                                echo "Container removal failed or was skipped"
                            }
                    }
                    }
                }
        stage('Build GOLANG Images') {
            steps {
                script {
                    def envContent = """
                        DB_USERNAME=${env.DB_USERNAME}
                        DB_PASSWORD=${env.DB_PASSWORD}
                        DB_NAME=${env.DB_NAME}
                        DB_HOST=${env.DB_HOST}${params.deployEnvironment}
                        DB_PORT=${env.DB_PORT}
                        ORIGIN=${env.ORIGIN}
                        ENV=${params.deployEnvironment}

                        ACCESS_TOKEN_PRIVATE_KEY=${env.ACCESS_TOKEN_PRIVATE_KEY}
                        REFRESH_TOKEN_PRIVATE_KEY=${env.REFRESH_TOKEN_PRIVATE_KEY}
                        ACCESS_TOKEN_EXPIRED_IN=${env.ACCESS_TOKEN_EXPIRED_IN}
                        REFRESH_TOKEN_EXPIRED_IN=${env.REFRESH_TOKEN_EXPIRED_IN}
                        GOOGLE_OAUTH_CLIENT_ID=${env.GOOGLE_OAUTH_CLIENT_ID}
                        GOOGLE_OAUTH_CLIENT_SECRET=${env.GOOGLE_OAUTH_CLIENT_SECRET}
                    """
                    writeFile file: '.env', text: envContent

                    // Display the content of the created .env file
                    echo "Content of .env:"
                    echo readFile('.env')
                    sh "echo ${params.deployEnvironment}"
                    sh "docker build -t  ${GOLANG_IMAGE_NAME}:${GIT_TAG} \
                    --build-arg DB_HOST=${env.DB_HOST}${params.deployEnvironment} \
                    --build-arg DB_USERNAME=${env.DB_USERNAME} \
                    --build-arg DB_NAME=${env.DB_NAME} \
                    --build-arg DB_PASSWORD=${env.DB_PASSWORD}\
                    --build-arg DB_PORT=${env.DB_PORT} \
                    --build-arg ENV=${params.deployEnvironment} \
                    --build-arg ORIGIN=${env.ORIGIN} ."
                }
            }
        }
        stage('Deploy') {
            steps {
                script {
                sh "docker run -d --name ${CONTAINER_NAME}-${params.deployEnvironment} --network ${params.deployEnvironment}-network ${GOLANG_IMAGE_NAME}:${GIT_TAG}"
                }
            }
        }

        stage('Clear Storage') {
            steps {
                script {
                    sh "docker image prune -a -f"
                }
            }
        }

        stage('Health Cheack') {
            steps {
                script {
                    def containerId = sh(script: "docker ps -q --filter name=${CONTAINER_NAME}-${params.deployEnvironment}", returnStdout: true)

                    if (containerId) {
                        def healthStatus = sh(script: "docker inspect --format '{{.State.Running}}'  ${containerId}", returnStdout: true)
                        
                        echo "Helath : ${healthStatus}"
                        if (healthStatus) {
                            echo "Container is running healthily."
                        } else {
                            error "Unable to retrieve container health status."
                        }
                    } else {
                        error "Container not found. Make sure it is running."
                    }
                }
            }
        }
    }

    post {
        success {
            echo 'Pipeline successfully completed!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}
