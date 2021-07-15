# Introduce controller

Some types of these controllers are:

1. Node controller: Responsible for noticing and responding when nodes go down.
2. Job controller: Watches for Job objects that represent one-off tasks, then creates Pods to run those tasks to completion.
3. Endpoints controller: Populates the Endpoints object (that is, joins Services & Pods).
4. Service Account & Token controllers: Create default accounts and API access tokens for new namespaces.

# cloud-controller-manager 
The following controllers can have cloud provider dependencies:

1. Node controller: For checking the cloud provider to determine if a node has been deleted in the cloud after it stops responding
2. Route controller: For setting up routes in the underlying cloud infrastructure
3. Service controller: For creating, updating and deleting cloud provider load balancers
