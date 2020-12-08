package tp2.sip.fixed2;

public abstract class Funcionario {
	
	protected double salarioBase;

	public abstract double getSalario();

	public double getSalarioBase() {
		return salarioBase;
	}

	public void setSalarioBase(double salarioBase) {
		this.salarioBase = salarioBase;
	}
}
