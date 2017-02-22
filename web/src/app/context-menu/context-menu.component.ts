import { Component, OnInit } from '@angular/core';
import { ContextMenuService } from '../context-menu.service';

@Component({
  selector: 'context-menu',
  templateUrl: './context-menu.component.html',
  styleUrls: ['./context-menu.component.css']
})
export class ContextMenuComponent implements OnInit {

	visible = false
	posX = '0px';
	posY = '0px';

  constructor(private menu: ContextMenuService) {
  	menu.visible.subscribe((visible) => this.visible = visible)
  	menu.positionX.subscribe((posX) => this.posX = `${posX}px`)
  	menu.positionY.subscribe((posY) => this.posY = `${posY}px`)
  }

  ngOnInit() {
  	
  }

}
