package tp2.dip.old;

public class Botao {
	
	private Lampada lampada;
	
	public void acionar() {
		if (lampada.isLigada()) {
			lampada.desligar();
		} else {
			lampada.ligar();
		}
	}
	
}
