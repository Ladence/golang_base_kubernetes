FROM scratch

ENV PORT 8080
EXPOSE $PORT

COPY base_service /
CMD "/base_service"