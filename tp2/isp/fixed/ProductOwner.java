package tp2.isp.fixed;

public class ProductOwner implements IProductOwner {

	private String nome;

	public String getNome(){
		return this.nome;
	}
	
	public void priorizarBacklog() {
		System.out.println("Priorizando backlog com base nas minhas necessidades de negócio");
	}

}