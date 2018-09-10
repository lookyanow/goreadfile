pipeline {
  agent any
  stages {
    stage('App build') {
      steps {
        sh 'go build '
      }
    }
    stage('Docker build') {
      steps {
        mail(subject: 'gohello build', body: 'gohello program builded successful', to: 'lookyanow@gmail.com')
      }
    }
  }
}