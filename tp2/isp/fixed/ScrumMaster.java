package tp2.isp.fixed;

public class ScrumMaster implements IScrumMaster {

	private String nome;

	public String getNome(){
		return this.nome;
	}
	
	public void blindarTime() {
		System.out.println("Devs trabalhando. Não se aproxime!");
	}

}