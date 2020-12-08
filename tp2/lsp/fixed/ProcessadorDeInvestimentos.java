package tp2.lsp.fixed;

import java.util.ArrayList;
import java.util.List;

public class ProcessadorDeInvestimentos {
	
	public static void main(String[] args) {
		
		for (ContaComum conta : contasDoBanco()) {
			conta.rende();
			
			System.out.println("Novo saldo:");
			System.out.println(conta.getSaldo());
		}
		
	}

	private static List<ContaComum> contasDoBanco() {
		return new ArrayList<ContaComum>();
	}

}
