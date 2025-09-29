# Fakejob Docker Image

This container runs a simple application that writes a few log messages to **stdout**, waits for a configurable amount of time, and then exits.  
You can control **how long it waits** and **whether it exits successfully or with an error** using environment variables.

## Docker Hub Repository

🐳 **Docker Hub**: https://hub.docker.com/repository/docker/rafaelbodao/fakejob

## Environment Variables

| Variable               | Description                                                                 | Default |
|------------------------|-----------------------------------------------------------------------------|---------|
| `SECONDS_TO_SLEEP`     | Number of seconds the application will sleep before finishing.               | `2`     |
| `SHOULD_FAIL`          | Set to `true` to force the application to exit with a non-zero (failure) code. Any other value or unset means success. | `false` |
| `SHOULD_FAIL_RANDOMLY` | Set to `true` to enable random failures. When enabled the application will randomly exit with a non-zero code (approximately 50% chance); any other value or unset means success. | `false` |

## Usage Examples

Run the container, sleeping for 10 seconds before exiting **successfully**:
```bash
docker run -e SECONDS_TO_SLEEP=10 rafaelbodao/fakejob
```

Run the container, sleeping for 10 seconds and then exiting with an error:
```bash
docker run -e SECONDS_TO_SLEEP=10 -e SHOULD_FAIL=true rafaelbodao/fakejob
```
