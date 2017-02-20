import { Injectable } from '@angular/core';
import { Http, Response, Headers, RequestOptions } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';

import { FileItem } from './file-item';

@Injectable()
export class FillyService {
  private api = '/browse'
  private headers = new Headers({ 'Content-Type': 'application/x-www-form-urlencoded' });
  private options = new RequestOptions({ headers: this.headers });

  constructor(private http: Http) { }

  getFiles(path: string): Observable<FileItem[]> {
  	return this.http.post(this.api, 'path='+path, this.options)
  		.map(this.fileMapper)
  }

  private fileMapper(response: Response): FileItem[] {
  	let data = response.json()
  	console.log("Data is", response.json())
  	return <FileItem[]>data
  }

}
