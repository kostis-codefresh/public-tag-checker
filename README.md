# Examples of Codefresh plugins in Golang

An example Codefresh plugin. It takes as argument a public dockerhub image (such httpd, ubuntu, alpine)
and returns its newest tag according to semantic versioning (completely unrelated to the `latest` tag

For the usage see [https://codefresh.io/steps/step/kostis-codefresh/newest-dockerhub-tag](https://codefresh.io/steps/step/kostis-codefresh%2Fnewest-dockerhub-tag)

The pipeline definition can also be found in [codefresh.yml](codefresh.yml).
