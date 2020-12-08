package tp2.isp.fixed;

public class Dev implements IDev {

	private String nome;

	public String getNome(){
		return this.nome;
	}
	
	public void implementarFuncionalidades() {
		System.out.println("Codando e tomando café compulsivamente!!");
	}

}