import { Component } from '@angular/core';
import { ContextMenuService } from './context-menu/context-menu.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  constructor(private menu: ContextMenuService) {

  }

  hideContextMenu(event) {
  	this.menu.hideMenu()
  }
}
