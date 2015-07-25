FROM scratch
ADD exile /bin/exile
ENTRYPOINT ["/bin/exile"]
