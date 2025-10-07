use actix_web::{App, HttpServer, Responder, web};
use broker_flattrade::Flattrade;
use core::Broker;

async fn place_order_handler() -> impl Responder {
    let flattrade = Flattrade::new();

    let flattrade_result = flattrade.place_order("RELIANCE", 10);

    format!("Flattrade order placed: {}", flattrade_result)
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    println!("Starting Actix server at http://localhost:8080");

    HttpServer::new(|| App::new().route("/place-order", web::get().to(place_order_handler)))
        .bind("127.0.0.1:8080")?
        .run()
        .await
}
