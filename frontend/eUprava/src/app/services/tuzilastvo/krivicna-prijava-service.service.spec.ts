import { TestBed } from '@angular/core/testing';

import { KrivicnaPrijavaServiceService } from './krivicna-prijava-service.service';

describe('KrivicnaPrijavaServiceService', () => {
  let service: KrivicnaPrijavaServiceService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(KrivicnaPrijavaServiceService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
