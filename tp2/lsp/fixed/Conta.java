package tp2.lsp.fixed;

public class Conta {
    
    protected double saldo;

    public Conta() {
		this.saldo = 0;
	}

    public void deposita(double valor) {
		this.saldo += valor;
	}

    public double getSaldo() {
		return saldo;
	}
}