import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PoterniceListComponent } from './poternice-list.component';

describe('PoterniceListComponent', () => {
  let component: PoterniceListComponent;
  let fixture: ComponentFixture<PoterniceListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PoterniceListComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PoterniceListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
