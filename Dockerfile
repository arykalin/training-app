FROM alpine
WORKDIR /bin
COPY training-app .
RUN chmod +x training-app
CMD ["./training-app"]