import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AppComponent } from './app.component';
import { FileListComponent } from './file-list/file-list.component';
import { FillyService } from './filly.service';

@NgModule({
  declarations: [
    AppComponent,
    FileListComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule
  ],
  providers: [HttpModule, FillyService],
  bootstrap: [AppComponent]
})
export class AppModule { }
