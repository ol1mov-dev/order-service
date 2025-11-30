<h1>Order Service</h1>

<p>Сервис заказов, где хранятся данные о заказах пользователя.</p>

<h3>Генерация protobuf файлов из  папки api/user_v1.</h3>
<pre>protoc --go_out=pkg --go_opt=paths=source_relative --go-grpc_out=pkg --go-grpc_opt=paths=source_relative api/user_v1/*.proto</pre>