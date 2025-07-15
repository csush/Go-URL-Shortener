# Go URL Shortener

## Scope
- Test-driven
- /POST URL to shorten, returns shortened URL
    - hash function that avoids collisions
- /GET/:shortenedURL, redirects to real URL
- /GET/all, returns all shortenedURL->realURL mappings
- In-memory DB to store data, inject as dependency so that can later switch with real DB
- Dockerfile
- CI/CD with Github
- Deployable to Render