pipeline {
    agent any

    parameters {
      choice(choices: ["dev", "sit", "prod"], description: "Which environment to deploy?", name: "deployEnvironment")
    }

    environment {
        GOLANG_IMAGE_NAME = "vocaverse-golang"
        IMAGE_TAG = "latest"
        CONTAINER_NAME = "vocaverse-golang"
    }

    stages {

        stage('Build DB Images') {
            steps {
                script {
                    sh "echo ${params.deployEnvironment}"
                    sh "docker build -t  ${GOLANG_IMAGE_NAME}:${IMAGE_TAG} \
                     --build-arg DB_HOST=${env.DB_HOST}${params.deployEnvironment} \
                     --build-arg DB_USER=${env.DB_USER} \
                     --build-arg DB_PASSWORD=${env.DB_PASSWORD}\
                     --build-arg DB_PORT=${env.DB_PORT} ."
                }
            }
        }
        stage ('Remove container'){
            steps {
              script {
                    // Run the command and capture the exit code
                    def exitCode = sh(script: "docker rm -f ${CONTAINER_NAME}-${params.deployEnvironment}", returnStatus: true)

                    // Check the exit code to determine success or failure
                    if (exitCode == 0) {
                        echo "Container removal was successful"
                        // Add more steps or logic here if needed
                    } else {
                        echo "Container removal failed or was skipped"
                        // Add more steps or logic here if needed
                    }
              }
            }
        }

        stage('Deploy') {
            steps {
                script {
                  sh "docker run -d --name ${CONTAINER_NAME}-${params.deployEnvironment} ${GOLANG_IMAGE_NAME}:${IMAGE_TAG}"
                }
            }
        }

        stage('Link Networks') {
            steps {
                script {

                  sh "docker network connect ${params.deployEnvironment}-network ${CONTAINER_NAME}-${params.deployEnvironment}"
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
