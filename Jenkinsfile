def service = 'pdf-print-server'
def version = ''
def dockerHub = 'eu.gcr.io/karnott-utilities'
pipeline {
    agent any
    environment {
      NEXUS_LOGIN = credentials('nexus-credentials')
    }
    options {
        disableConcurrentBuilds()
    }
    stages {
      stage('Project Build') {
      when {
          beforeAgent true
          anyOf {
            not { changelog '.*\\[ci skip\\].*' }
            not { branch 'develop' }
          }
      }
      stages {
        stage('build stage') {
          steps {
            sh "docker build -t ${service} . -f Dockerfile.build"
          }
        }
        stage('build & push develop stage'){
          when {
            branch 'develop'
          }
          steps {
             // create a pre release in develop branch
            sh "docker run -v ${workspace}:/usr/src/app --rm ${dockerHub}/release-manager:1.0.1"

            script {
              version = sh(returnStdout: true, script: "cat VERSION").trim()
            }
            sh "docker build -t ${service} . -f Dockerfile.build"
            sh "docker tag ${service} ${dockerHub}/${service}:${version}"
            sh "docker push ${dockerHub}/${service}:${version}"
            sh "docker tag ${service} ${dockerHub}/${service}:latest"
            sh "docker push ${dockerHub}/${service}:latest"

            withCredentials([usernamePassword(credentialsId: 'github_plain_auth', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]) {
              sh "git remote -v"
              sh "git push https://${GIT_USERNAME}:${GIT_PASSWORD}@github.com/Karnott/${service}.git --follow-tags HEAD:develop"
              sh "git push https://${GIT_USERNAME}:${GIT_PASSWORD}@github.com/Karnott/${service}.git --tags"
            }
          }
        }

        stage('build & push production stage'){
          when {
            branch 'master'
          }
          steps {
            script {
              version = sh(returnStdout: true, script: "cat VERSION").trim()
            }

            sh "docker build -t ${service} . -f Dockerfile.build"
            sh "docker tag ${service} ${dockerHub}/${service}:lts"
            sh "docker push ${dockerHub}/${service}:lts"
          }
        }

        stage('Deploy') {
          when {
            anyOf {
              branch 'master'
              branch 'develop'
            }
          }
          steps {
            script {
              service_env = 'recette'
              if(env.BRANCH_NAME == 'master') {
                service_env = 'production'
              }
              build job: "infrastructure/update", parameters: [string(name: 'ENVIRONMENT', value: service_env), string(name: 'SERVICE_VERSION', value: "${version}"), string(name: 'SERVICE_NAME', value: "${service}")]
            }
          }
        }
      }
    }
  }

  post {
    always{
      cleanWs()
    }
  }
}
