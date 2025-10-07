pub trait Broker {
    fn place_order(&self, symbol: &str, quantity: u32) -> bool;
}
