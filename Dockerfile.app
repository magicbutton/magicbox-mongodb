FROM mcr.microsoft.com/powershell
# misses the az cli 

RUN pwsh -c "Install-Module -Name PnP.PowerShell -Force -AllowPrerelease -Scope AllUsers;" 

RUN apt update -y
RUN apt upgrade -y
RUN apt install golang-1.21 -y
ENV GOBIN="/usr/local/bin"
ENV PATH="/usr/lib/go-1.21/bin:${PATH}"

ENV KITCHEN_HOME="/kitchens"

WORKDIR /kitchens/magicbox-mongodb
# The go program can have dependencies to files in the project root
# so that is why we copy the entire project to the container


RUN go install github.com/koksmat-com/koksmat@v2.0.0.8

COPY . .  
WORKDIR /kitchens/magicbox-mongodb/.koksmat/app

RUN go install

RUN koksmat context init sharepoint


CMD [ "sleep","infinity"]