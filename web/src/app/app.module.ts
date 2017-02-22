import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AppComponent } from './app.component';
import { FileListComponent } from './file-list/file-list.component';
import { FillyService } from './filly.service';
import { ContextMenuService } from './context-menu.service';
import { ContextMenuComponent } from './context-menu/context-menu.component';

@NgModule({
  declarations: [
    AppComponent,
    FileListComponent,
    ContextMenuComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule
  ],
  providers: [HttpModule, FillyService, ContextMenuService],
  bootstrap: [AppComponent]
})
export class AppModule { }
