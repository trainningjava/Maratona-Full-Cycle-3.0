FROM gobuffalo/buffalo:latest 
RUN go get -u github.com/gobuffalo/buffalo/buffalo
WORKDIR /src/
RUN buffalo new --api --skip-pop desafio2
RUN echo "C O P Y N G . . . . "
WORKDIR /src/desafio2
COPY . /src/
RUN echo "======================"
RUN echo "A C T I O N S . . . . "
RUN echo "======================"
RUN ls -la /src/desafio2/actions
RUN echo "======================"
RUN echo "T E M P L A T E S . . "
RUN echo "======================"
RUN ls -la /src/desafio2/templates
RUN npm install
RUN buffalo build -o /bin/app

# Bind the app to 0.0.0.0 so it can be seen from outside the container
ENV ADDR=0.0.0.0
EXPOSE 3000

# Uncomment to run the migrations before running the binary:
# CMD /bin/app migrate; /bin/app
CMD ["/bin/app"]