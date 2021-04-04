/*/

===============================================================================
=  This file is to create a nginx config for VPN proxy

server {
    server_name vpn.example.com;
    listen [::]:443  default ipv6only=off ssl;

    ssl on;
    ssl_certificate      /etc/lets.encrypt/ssl/vpn.example.com.full;
    ssl_certificate_key  /etc/lets.encrypt/ssl/vpn.example.com.key;
    ssl_dhparam          /etc/nginx/ssl/dhparam.pem;

    ssl_prefer_server_ciphers on;
    ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
    ssl_ciphers 'ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-DSS-AES128-GCM-SHA256:kEDH+AESGCM:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-DSS-AES128-SHA256:DHE-RSA-AES256-SHA256:DHE-DSS-AES256-SHA:DHE-RSA-AES256-SHA:ECDHE-RSA-DES-CBC3-SHA:ECDHE-ECDSA-DES-CBC3-SHA:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:AES:CAMELLIA:DES-CBC3-SHA:!aNULL:!eNULL:!EXPORT:!DES:!RC4:!MD5:!PSK:!aECDH:!EDH-DSS-DES-CBC3-SHA:!EDH-RSA-DES-CBC3-SHA:!KRB5-DES-CBC3-SHA';
    add_header Strict-Transport-Security "max-age=31536000";

    proxy_buffering    off;
    proxy_buffer_size  128k;
    proxy_buffers 100  128k;

    ## VPN server ..
    location /vpn {

       proxy_set_header      X-Forwarded-For $remote_addr;
       proxy_pass            http://127.0.0.1:9000;
       proxy_http_version    1.1;
       proxy_set_header      Upgrade $http_upgrade;
       proxy_set_header      Connection "upgrade";
       proxy_read_timeout    86400;
       proxy_connect_timeout 43200000;

       tcp_nodelay on;
   }
}


/*/