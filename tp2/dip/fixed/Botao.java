package tp2.dip.fixed;

public class Botao {
	
	public void acionar(Dispositivo dispositivo) {

		if (dispositivo.isOn()) {
			dispositivo.desligar();
		} else {
			dispositivo.ligar();
		}
	}
	
}
