import { TestBed } from '@angular/core/testing';

import { NgxAstilectronService } from './ngx-astilectron.service';

describe('NgxAstilectronService', () => {
  let service: NgxAstilectronService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(NgxAstilectronService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
