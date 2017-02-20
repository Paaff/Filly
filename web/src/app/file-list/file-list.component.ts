import { Component, OnInit } from '@angular/core';
import { HttpModule } from '@angular/http';
import { FillyService } from '../filly.service';
import { File } from '../file';

@Component({
  selector: 'file-list',
  templateUrl: './file-list.component.html',
  styleUrls: ['./file-list.component.css']
})
export class FileListComponent implements OnInit {

  files: File[]

  constructor(private http: HttpModule, private filly: FillyService) { }

  ngOnInit() {
  	this.getFiles('')
  }

  getFiles(path: string) {
  	this.filly.getFiles(path).subscribe(
  		files => this.files = files
  	)
  }

}
