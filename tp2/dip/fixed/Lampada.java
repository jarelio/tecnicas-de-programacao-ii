package tp2.dip.fixed;

public class Lampada implements Dispositivo{
	
	private boolean on = false;

	public boolean isOn() {
		return this.on;
	}
	
	public void ligar() {
		System.out.println("Lâmpada ligada");
	}

	public void desligar() {
		System.out.println("Lâmpada desligada");
	}

}
