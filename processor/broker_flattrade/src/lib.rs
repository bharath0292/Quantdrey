use core::Broker;

pub struct Flattrade;

impl Broker for Flattrade {
    fn place_order(&self, symbol: &str, quantity: u32) -> bool {
        println!("Flattrade placing order for {} x {}", symbol, quantity);
        true
    }
}

impl Flattrade {
    pub fn new() -> Self {
        Flattrade
    }
}
