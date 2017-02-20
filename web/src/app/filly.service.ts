import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/map';

import { File } from './file';

@Injectable()
export class FillyService {
  private api = '/browse'

  constructor(private http: Http) { }

  getFiles(path: string): Observable<File[]> {
  	return this.http.post(this.api, 'path='+path)
  		.map((r: Response) => <File[]>r.json().data)
  }

}
