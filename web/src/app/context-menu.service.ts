import { Injectable } from '@angular/core';
import { Subject }    from 'rxjs/Subject';

@Injectable()
export class ContextMenuService {

  constructor() { }

  visible = new Subject<boolean>();
  positionX = new Subject<number>();
  positionY = new Subject<number>();

  showMenu(X: number, Y: number) {
  	console.log("Showing menu at", X, Y)
  	this.positionX.next(X)
  	this.positionY.next(Y)
  	this.visible.next(true)
  }

  hideMenu() {
  	this.visible.next(false)
  }

}
