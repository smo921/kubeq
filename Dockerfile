FROM scratch
MAINTAINER Stephen Oberther <stephenoberther@gmail.com>
COPY kubeq /kubeq
CMD ["/kubeq"]
