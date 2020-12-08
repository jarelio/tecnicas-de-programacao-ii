package tp2.sip.fixed;

import static tp2.sip.fixed.Cargo.DBA;
import static tp2.sip.fixed.Cargo.DESENVOLVEDOR;
import static tp2.sip.fixed.Cargo.GERENTE;

public class Funcionario {
	
	private Cargo cargo;
	private double salarioBase;

	public Cargo getCargo() {
		return cargo;
	}

	public void setCargo(Cargo cargo) {
		this.cargo = cargo;
	}

	public double getSalarioBase() {
		return salarioBase;
	}

	public void setSalarioBase(double salarioBase) {
		this.salarioBase = salarioBase;
	}

	public double getSalario() {
		if(DESENVOLVEDOR.equals(this.cargo)) {
			return this.salarioBase * 1.1;
		}

		if(DBA.equals(this.cargo)) {
			return this.salarioBase * 1.2;
		}
		
		if(GERENTE.equals(this.cargo)) {
			return this.salarioBase * 1.3;
		}
		
		throw new RuntimeException("funcionário inválido");
	}
}
