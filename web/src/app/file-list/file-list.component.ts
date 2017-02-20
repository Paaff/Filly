import { Component, OnInit } from '@angular/core';
import { HttpModule } from '@angular/http';
import { FillyService } from '../filly.service';
import { FileItem } from '../file-item';

@Component({
  selector: 'file-list',
  templateUrl: './file-list.component.html',
  styleUrls: ['./file-list.component.css']
})
export class FileListComponent implements OnInit {

  path: string[] = [];
  files: FileItem[];

  constructor(private http: HttpModule, private filly: FillyService) { }

  ngOnInit() {
  	this.getFiles('')
  }

  test() {
  	console.log(this.files)
  }

  gotoFolder(file?: FileItem) {
  	if (file && file.type != 'folder') return
  	if (file) {
      let folder = file.path.split(/[\\/]/).pop()
  	  this.path.push(folder)
  	}
  	console.log("Going to folder", this.path.join('\\'))
  	this.filly.getFiles(this.path.join('\\')).subscribe(
  		files => this.files = files,
  		this.onError
  	)
  }

  back() {
  	this.path.pop()
  	this.gotoFolder()
  }

  getFiles(path: string) {
  	this.filly.getFiles(path).subscribe(
  		files => this.files = files,
  		this.onError
  	)
  }

  private onError(err) {
  	console.error(err)
  }

}
