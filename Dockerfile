FROM scratch
MAINTAINER Stephen Oberther <stephenoberther@gmail.com>
COPY templates/ /templates/
COPY kubeq.conf /kubeq.conf
COPY kubeq /kubeq
CMD ["/kubeq"]
