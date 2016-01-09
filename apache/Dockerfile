FROM ubuntu:14.04

RUN apt-get update  
RUN apt-get -y install \
        curl \
        wget \
        vim \
        apache2 \
        libapache2-mod-php5 \
        php5 \
        php5-mongo \
        php5-gd \
        php5-curl \
        php5-imagick \
        php5-tidy 

RUN a2enmod php5
RUN a2enmod rewrite

RUN apt-get clean
RUN apt-get autoclean
RUN rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* 


RUN echo "ServerName localhost" >> /etc/apache2/apache2.conf \
    && sed -i 's/variables_order.*/variables_order = "EGPCS"/g' \
        /etc/php5/apache2/php.ini

ADD common.ini /etc/php5/apache2/conf.d/
ADD common.ini /etc/php5/cli/conf.d/

RUN mkdir -p /app && rm -rf /var/www/html 
COPY . /app  
WORKDIR /app  
RUN chmod 755 ./start.sh

RUN mkdir -p /site && ln -s /site/ /var/www/html

VOLUME ["/var/log/apache2","/site"]
EXPOSE 80  
CMD ["./start.sh"]  
