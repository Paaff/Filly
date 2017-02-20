/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { FillyService } from './filly.service';

describe('FillyService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [FillyService]
    });
  });

  it('should ...', inject([FillyService], (service: FillyService) => {
    expect(service).toBeTruthy();
  }));
});
