import { ComponentFixture, TestBed } from '@angular/core/testing';

import { KListComponent } from './k-list.component';

describe('KListComponent', () => {
  let component: KListComponent;
  let fixture: ComponentFixture<KListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ KListComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(KListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
