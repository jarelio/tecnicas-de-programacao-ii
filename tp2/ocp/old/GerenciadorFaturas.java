package tp2.ocp.old;

public class GerenciadorFaturas {
	
	private FilaContabilidade filaContabilidade;
	private EnviadorEmail enviadorEmail;
	
	public void pagarFatura(Fatura fatura) {
		fatura.setPaga(true);
		
		filaContabilidade.enviarParaFilaProcessamento(fatura);
		enviadorEmail.enviar("email@email.com");
	}

}
