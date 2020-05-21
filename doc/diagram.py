from diagrams import Diagram
from diagrams import Cluster
from diagrams.gcp.network import LoadBalancing, DNS
from diagrams.gcp.database import BigTable,Memorystore
from diagrams.gcp.storage import Storage
from diagrams.gcp.compute import ComputeEngine
from diagrams.onprem.monitoring import Prometheus, Grafana


with Diagram("Pastebin architecture",show=False):
    dns = DNS("dns")
    elb = LoadBalancing("loadbalancer")
    with Cluster("REST"):
        gces = [ComputeEngine("gce") for  _  in range(3)]
        cache_m = Memorystore("cache")

    with Cluster("NoSQL"):
        db = BigTable("bigtable1")
    store = Storage("filestore")
    with Cluster("monitoring"):
        prom = Prometheus("prometheus")
        graf = Grafana("grafana")

    dns >> elb 
    elb >> gces
    elb >> graf
    gces >> cache_m >> db
    gces >>  store >> gces
    gces - prom
    graf >> prom