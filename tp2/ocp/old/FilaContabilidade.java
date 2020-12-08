package tp2.ocp.old;

public class FilaContabilidade {
	
	public void enviarParaFilaProcessamento(Fatura fatura) {
		System.out.println("Enviando fatura " + fatura.getNumero() + " para a fila de processamento...");
	}

}
