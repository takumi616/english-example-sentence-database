FROM postgres:latest

RUN apt-get update \
    && apt-get install -y --no-install-recommends wget ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Install psqldef
RUN wget --no-check-certificate https://github.com/sqldef/sqldef/releases/download/v0.16.15/psqldef_linux_amd64.tar.gz \
    && tar xf psqldef_linux_amd64.tar.gz \
    && mv psqldef /usr/local/bin/psqldef \
    && rm psqldef_linux_amd64.tar.gz \
    && chmod +x /usr/local/bin/psqldef