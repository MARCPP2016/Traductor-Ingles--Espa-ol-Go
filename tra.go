/*
************************
Traductor Ingles-Español
(Marcpp)
Montevideo Uruguay
noviembre-2016
**********************
para foro cristallab
*/

/*
***************************************************
Bueno un pequeño aporte,un traductor en Go, recordar
que el codigo hay que depurarlo, ya que seguramente
al contener mas de 500 lineas!  contendra errores.
**************************************************
*/

package main

import "fmt" 

//Funcion principal
func main() {
	
//Llamamos a funcion que contiene codigo.
tra()

}

/*
********************************************************
Recordar tambien , que al contener muchas lineas de codigo
se recomienda,subdividir en modulos.
Esto en caso de depuarcion y mejora del codigo ,facilita
el trabajo de un tercero.

Algo mas, veran en el codigo correspondiente a la sentencia
switch que en el final de cada variable existe una(;)
y se preguntaran¡???.la razon es que importe lineas de
codigo desde otros lenguajes de programacion
que yo mismo estoy programando.
Asi que no se preocupen ,si quitaran las(;)
compilara igual. 
**********************************************************
*/

//Funcion tra
func tra(){

//Bucle infinito
for true {
var y string

fmt.Println("\t\tColoca tu palabra")
fmt.Println("\t\t(Crl_c salir)\n")
fmt.Scanf("%s\n",&y)

//Sentencia switch
switch(y){
	 	
	 	//letra a
		case "aqua" :fmt.Println(": agua");
        break;
	    case "abril" :fmt.Println(": abril");
		break;
		case "august" :fmt.Println(": agosto");
		break;
		case "alone" :fmt.Println(": solo,solitario");
		break;
		case "aback":fmt.Println(": take aback");
		break;
		case "abacus":fmt.Println(": abaco");
		break;
		case "abandon":fmt.Println(": abandonar");
		break;
		case "abase":fmt.Println(": humillar");
		break;
	    case "abashed":fmt.Println(": confundido");
		break;
		case "abate":fmt.Println(": confundido");
		break;
		case "abattoir" :fmt.Println(": matadero");
		break;
        case "abbess" :fmt.Println(": abadesa");
		break;
        case "abbey" :fmt.Println(": abadia");
		break;
        case "abbot" :fmt.Println(": abad");
		break;
        case "abbreviate" :fmt.Println(": abreviatura");
		break;
        case "ABC" :fmt.Println(": abecedario");
		break;
        case "abdicate" :fmt.Println(": abdicar");
		break;
        case "abdication" :fmt.Println(": abdicacion");
		break;
        case "abdomen" :fmt.Println(": abdomen");
		break;
        case "abduct" :fmt.Println(": raptar");
		break;
        case "abduction" :fmt.Println(": rapto,secuestro");
		break;
        case "abed" :fmt.Println(": en cama");
		break;
        case "aberration" :fmt.Println(": aberracion");
		break;
        case "abet" :fmt.Println(": apoyar un delito");
		break;
        case "abettor" :fmt.Println(": complice");
		break;
        case "abeyance" :fmt.Println(": suspension");
		break;
        case "abhor" :fmt.Println(": aborrecer");
		break;
        case "abide" :fmt.Println(": cumplir con");
		break;
        case "abiding" :fmt.Println(": constante,duradero");
		break;
        case "ability" :fmt.Println(": abilidad,aptitud");
		break;
        case "abject" :fmt.Println(": abyecto");
		break;
        case "abjection" :fmt.Println(": abyeccion");
		break;
        case "abjure" :fmt.Println(": abjuntar");
		break;
        case "albative" :fmt.Println(": ablativo");
		break;
        case "ablaze" :fmt.Println(": en llamas");
		break;
        case "able" :fmt.Println(": habil,capaz");
		break;
        case "ablution" :fmt.Println(": ablucion");
		break;
        case "ably" :fmt.Println(": habilmente");
		break;
        case "abnegation" :fmt.Println(": abnegacion");
		break;
        case "abnormal" :fmt.Println(": anormal");
		break;
        case "aboard" :fmt.Println(": a bordo");
		break;
		case "abolish" :fmt.Println(": abolir");
		break;
      	case "abolition" :fmt.Println(": abolicion");
		break;
      	case "A-bomob" :fmt.Println(": bomba atomica");
		break;
    	case "abominable" :fmt.Println(": abominable");
		break;
        case "abominate" :fmt.Println(": abominar");
		break;
        case "abomination" :fmt.Println(": abominacion");
		break;
        case "aboriginal" :fmt.Println(": aborigen");
		break;
        case "aborigines" :fmt.Println(": aborigenes");
		break;
        case "abort" :fmt.Println(": abortar");
		break;
        case "abortion" :fmt.Println(": aborto");
		break;
        case "abortionist" :fmt.Println(": partidario del aborto");
		break;
        case "abound" :fmt.Println(": abundar");
		break;
        case "about" :fmt.Println(": acerca");
		break;
        case "above" :fmt.Println(": sobre");
		break;
		case "aboveboard" :fmt.Println(": franco,abiertamente");
		break;
		case "abracadabra" :fmt.Println(": abracadabra");
		break;
	    case "abrasion" :fmt.Println(": abrasion");
		break;
        case "abrasive" :fmt.Println(": abrasivo");
		break;
        case "abreast" :fmt.Println(": lado a lado");
		break;
        case "abridge" :fmt.Println(": abreviar,condensar");
		break;
		case "apple" :fmt.Println(": manzana");
		break;
		case "after" :fmt.Println(": despues");
		break;
    
		//letra b
		case "blue" :fmt.Println(": azul");
		break;
		case "before" :fmt.Println(": antes");
		break;
        case "bread" :fmt.Println(": pan");
		break;
		case "breath" :fmt.Println(": aliento,respiracion");
		break;


		case "bitch" :fmt.Println(": puta,perra");
		break;
		
		case "boyfriend" :fmt.Println(": novio");
		break;
		case "babe" :fmt.Println(": bebe");
		break;
        case "boy" :fmt.Println(": chico");
		break;
        case "brother" :fmt.Println(": hermano");
		break;
        case "book" :fmt.Println(": libro");
        break;
        case "bike" :fmt.Println(": bicicleta");
		break;
        case "boat" :fmt.Println(": bote,lancha");
		break;
        case "black" :fmt.Println(": negro");
		break;
		case "bee" :fmt.Println(": abeja");
		break;
		case "blade" :fmt.Println(": espada, cuchilla,hoja");
		break;
        //letra c
		case "car" :fmt.Println(": auto");
		break;
        case "chilhood" :fmt.Println(": niñez,infancia");
		break;
		case "child" :fmt.Println(": niño");
		break;
		case "climup" :fmt.Println(": trepar");
		break;
		case "crawling" :fmt.Println(": agacharse,rebajarse,arrastrarse");
		break;
		case "crunch" :fmt.Println(": crujido");
		break;
		case "country" :fmt.Println(": pais");
		break;
        case "city" :fmt.Println(": ciudad");
		break;
        case "clumsy" :fmt.Println(": torpe");
		break;
        case "chair" :fmt.Println(": silla");
		break;
		case "cat" :fmt.Println(": gato");
		break;
		case "climp" :fmt.Println(": subir");
		break;
		case "cementery" :fmt.Println(": cementerio");
		break;
		
		
		//letra d
		case "dog" :fmt.Println(": perro");
		break;
		case "dead" :fmt.Println(": muerte");
		break;
		case "down" :fmt.Println(": abajo");
		break;
		case "december" :fmt.Println(": diciembre");
		break;
	
        //letra e
        case "eat" :fmt.Println(": comer");
		break;
	    case "eight" :fmt.Println(": ocho");
		break;
	    
        case "electricity" :fmt.Println(": electricidad");
		break;
		case "eternal" :fmt.Println(": eterno");
		break;
		case "eternity" :fmt.Println(": eternidad");
		break;
		case "eleven" :fmt.Println(": once");
		break;
	    
		//letra f
		case "four" :fmt.Println(": cuatro");
		break;
		case "five" :fmt.Println(": cinco");
		break;
	    case "February" :fmt.Println(": febrero");
		break;
		case "friday" :fmt.Println(": viernes");
		break;
		case "fight" :fmt.Println(": pelea");
		break;
		case "fly" :fmt.Println(": volar/mosca");
		break;
        case "father" :fmt.Println(": padre");
		break;
		case "friend" :fmt.Println(": amigo");
		break;
        case "fire" :fmt.Println(": fuego");
		break;
        	
		//letra g
		case "game" :fmt.Println(": juego");
		break;
		case "grass" :fmt.Println(": hierva,gramilla,cesped");
		break;
		case "girl" :fmt.Println(": chica");
		break;
        case "girlfriend" :fmt.Println(": novia");
		break;
        case "gentleman" :fmt.Println(": señor,caballero");
		break;
		case "guest" :fmt.Println(": invitado");
		break;
	    case "garden" :fmt.Println(": jardin");
		break;
		
		case "good" :fmt.Println(": bien,bueno");
		break;
		case "god" :fmt.Println(": Dios");
		break;
		case "glass" :fmt.Println(": vidrio espejo");
		break;
		case "gun" :fmt.Println(": arma");
		break;
        //letra h
        case "refrigerator" :fmt.Println(": heladera");
		break;
		case "horse" :fmt.Println(": caballo");
		break;
        
        case "hard" :fmt.Println(": duro");
		break;
		case "house" :fmt.Println(": casa");
		break;
		case "happy" :fmt.Println(": feliz");
		break;
		case "heaven" :fmt.Println(": cielo,(sagrado)");
		break;
		
		//letra i
		case "imortal" :fmt.Println(":imortal");
		break;
		case "inn" :fmt.Println(":hotel,alojamiento");
		break;
		case "ice" :fmt.Println(": hielo");
		break;
	    
		
		//letra j
		case "june" :fmt.Println(":junio");
		break;
	    case "july" :fmt.Println("julio");
		break;
	
		case "january" :fmt.Println(":enero");
		break;
		case "jump" :fmt.Println(":saltar");
		break;
		case "junk" :fmt.Println(":basura,chatarra,hierro viejo");
		break;
		case "joy" :fmt.Println(":alegria, gozo,disfrutar");
		break;

		//letra k
		case "know" :fmt.Println(":saber,conocer");
		break;
		case "knife" :fmt.Println(":cuchillo");
		break;
		case "knight" :fmt.Println(":caballero");
		break;
	    //letra l
		case "lemon" :fmt.Println(":limon");
		break;
		case "leave" :fmt.Println(":partir");
		break;
		case "lost" :fmt.Println(":perdido");
		break;
		case "lose" :fmt.Println(":perder");
		break;
		case "loose" :fmt.Println(":floja");
		break;
		case "life" :fmt.Println(":vida");
		break;
		case "left" :fmt.Println(":izquierda");
		break;
		case "long" :fmt.Println(":largo");
		break;
		case "last" :fmt.Println(":ultimo");
		break;
		case "light" :fmt.Println(":luz");
		break;
		case "later" :fmt.Println(":mas tarde");
		break;
		case "late" :fmt.Println(":tarde");
		break;
        //letra m
        case "music" :fmt.Println(": musica");
		break;
	
	    case "march" :fmt.Println(":marzo");
		break;
	    case "match" :fmt.Println(":partido,juego");
		break;
		case "monday" :fmt.Println(":lunes");
		break;
		case "meet" :fmt.Println(":conocerse");
		break;
		case "meat" :fmt.Println(":carne");
		break;
		case "mother" :fmt.Println(":madre");
		break;
		case "mathematics" :fmt.Println(":matematicas");
		break;
		case "man" :fmt.Println(":hombre");
		break;
		case "may" :fmt.Println(":mayo");
		break;
	    case "metal" :fmt.Println(": metal");
		break;
	
		//letra n
		case "nine" :fmt.Println(":nueve");
		break;
	    
		case "nail" :fmt.Println(":clavo");
		break;
        case "nephew" :fmt.Println(":sobrino");
		break;
		case "niece" :fmt.Println(":sobrina");
		break;
		case "newspaper" :fmt.Println(":periodico");
		break;
		case "november" :fmt.Println(":noviembre");
		break;
	
        //letra ñ
		//letra o
		case "october" :fmt.Println(":octubre");
		break;
		case "one" :fmt.Println(":uno");
		break;
	
	
		case "oil" :fmt.Println(":oleo, nafta,petroleo");
		break;
		case "orange" :fmt.Println(":naranja");
		break;
		
		//letra p
		case "paper" :fmt.Println(":papel");
		break;
        case "pool" :fmt.Println(":picina");
		break;
		case "punk" :fmt.Println(":pobre hombre,ladron ,mala gente");
		break;
	    case "plate" :fmt.Println(":plato");
		break;
        case "park" :fmt.Println(":parque");
		break;
		case "play" :fmt.Println(":tocar,jugar");
		break;
		
	    //letra q
	   	case "quest" :fmt.Println(":busqueda");
		break;
	
		//letra r
		case "right" :fmt.Println(":derecha");
		break;
		case "red" :fmt.Println(":rojo");
		break;
	    case "run" :fmt.Println(":correr");
		break;
		case "rain" :fmt.Println(":lluvia");
		break;
	    case "round" :fmt.Println(": redondo");
		break;
		//letra s
		case "shop" :fmt.Println(": comprar");
		break;
		case "sea" :fmt.Println(": mar,costa");
		break;
	    case "sell" :fmt.Println(": venta");
		break;
	    case "six" :fmt.Println(": seis");
		break;
	    case "seven" :fmt.Println(": siete");
		break;
	    case "september" :fmt.Println(": septiembre");
		break;
		case "sunday" :fmt.Println(": domingo");
		break;
		case "saturday" :fmt.Println(": sabado");
		break;
		case "stairs" :fmt.Println(": escalera");
		break;
        case "ladder" :fmt.Println(": escalera");
		break;
        case "stars" :fmt.Println(": estrella");
		break;
        case "spanish" :fmt.Println(": español");
		break;
     	case "sunshine" :fmt.Println(": soleado");
		break;
        case "smoke" :fmt.Println(": humo");
		break;
		case "short" :fmt.Println(": corto");
		break;
		case "switch" :fmt.Println(": cambio,intercambio");
		break;
		case "sister" :fmt.Println(": hermana");
		break;
		case "shadow" :fmt.Println(": sombra");
		break;
		case "shoes" :fmt.Println(": zapatos");
		break;
		case "street" :fmt.Println(": calle");
		break;
		case "sex" :fmt.Println(": sexo");
		break;
        case "son" :fmt.Println(": hijo");
		break;
		case "sad" :fmt.Println(": triste");
		break;
		case "square" :fmt.Println(": cuadrado");
		break;
		case "sphere" :fmt.Println(": esfera");
		break;
		case "space" :fmt.Println(": espacio");
		break;
		case "sky" :fmt.Println(": cielo");
		break;
		case "sleep" :fmt.Println(": dormir");
		break;
		//letra t
		case "ten" :fmt.Println(": diez");
		break;
	    case "two" :fmt.Println(": dos");
		break;
	    case "three" :fmt.Println(": tres");
		break;
	    case "Thursday" :fmt.Println(": jueves");
		break;
        case "tuesday" :fmt.Println(": martes");
		break;
        case "town" :fmt.Println(": pueblo");
		break;
        case "table" :fmt.Println(": mesa");
		break;
	    case "time" :fmt.Println(": tiempo");
		break;
		case "timer" :fmt.Println(": cronometro");
		break;
		case "tall" :fmt.Println(": alto");
		break;
		case "tank" :fmt.Println(": tanque");
		break;
		case "triangle" :fmt.Println(": triangulo");
		break;
		//letra u
		case "uncle" :fmt.Println(": tio");
		break;
        case "up" :fmt.Println(": arriba");
		break;
		case "unique" :fmt.Println(": unico");
		break;
		case "ugly" :fmt.Println(": horrible");
		break;
		//letra v
		case "vein" :fmt.Println(": vena");
		break;
		case "variable" :fmt.Println(": variable");
		break;
	    //letra w
	    case "woman" :fmt.Println(": mujer");
		break;
	    case "water" :fmt.Println(": agua");
		break;
	    case "while" :fmt.Println(": mientras");
		break;
		case "wind" :fmt.Println(": viento");
		break;
		case "wild" :fmt.Println(": salvaje");
		break;
		case "winter" :fmt.Println(": invierno");
		break;
		case "win" :fmt.Println(": ganar");
		break;
		case "wine" :fmt.Println(": vino");
		break;
		case "winer" :fmt.Println(": ganador");
		break;
		case "wallet" :fmt.Println(": cartera,monedero,billetera");
		break;
		case "wise" :fmt.Println(": sabio");
		break;
		case "Wednesday": fmt.Println(": miercoles");
		break;
		case "War": fmt.Println(": guerra");
		break;
        case "Wood": fmt.Println(": madera");
		break;
        case "Woods": fmt.Println(": bosque");
		break;
		//letra x
		//letra y
		case "yellow" :fmt.Println(": amarillo");
		break;
		//letra z
		case "zoo" :fmt.Println(": zoologico");
		break;
	

    	}
       	  
   }
}





