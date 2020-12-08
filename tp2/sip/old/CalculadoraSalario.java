package tp2.sip.old;

import static tp2.sip.old.Cargo.DBA;
import static tp2.sip.old.Cargo.DESENVOLVEDOR;
import static tp2.sip.old.Cargo.GERENTE;

public class CalculadoraSalario {
	
	public double calcularSalario(Funcionario funcionario) {
		if(DBA.equals(funcionario.getCargo())) {
			return funcionario.getSalarioBase() * 1.2;
		}
		
		if(GERENTE.equals(funcionario.getCargo())) {
			return funcionario.getSalarioBase() * 1.3;
		}
		
		if(DESENVOLVEDOR.equals(funcionario.getCargo())) {
			return funcionario.getSalarioBase() * 1.1;
		}
		
		throw new RuntimeException("funcionário inválido");
	}

}
