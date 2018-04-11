import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'unicardf'
})
export class UnicardfPipe implements PipeTransform {
  transform(value: number): string {
  	var uoffset = 126976;
    var valv = value % 13;
    var suitv = value / 13;
    switch (suitv) {
    	case Suit.Spades:
    		suitv = 10*16;
    		break;
    	case Suit.Hearts:
    		suitv = 11*16;
    		break;
    	case Suit.Diamonds:
    		suitv = 12*16;
    		break;
    	case Suit.Clubs:
    		suitv = 13*16;
    		break;
		default:break;
    }
    var cardv = (uoffset + suitv + valv);
    return String.fromCodePoint(cardv);
  }

}
